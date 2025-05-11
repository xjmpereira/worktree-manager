const std = @import("std");

// Although this function looks imperative, note that its job is to
// declaratively construct a build graph that will be executed by an external
// runner.
pub fn build(b: *std.Build) void {
    // Standard target options allows the person running `zig build` to choose
    // what target to build for. Here we do not override the defaults, which
    // means any target is allowed, and the default is native. Other options
    // for restricting supported target set are available.
    const target = b.standardTargetOptions(.{});

    // Standard optimization options allow the person running `zig build` to select
    // between Debug, ReleaseSafe, ReleaseFast, and ReleaseSmall. Here we do not
    // set a preferred release mode, allowing the user to decide how to optimize.
    const optimize = b.standardOptimizeOption(.{});

    // This creates a "module", which represents a collection of source files alongside
    // some compilation options, such as optimization mode and linked system libraries.
    // Every executable or library we compile will be based on one or more modules.
    const lib_mod = b.createModule(.{
        // `root_source_file` is the Zig "entry point" of the module. If a module
        // only contains e.g. external object files, you can make this `null`.
        // In this case the main source file is merely a path, however, in more
        // complicated build scripts, this could be a generated file.
        .root_source_file = b.path("src/root.zig"),
        .target = target,
        .optimize = optimize,
    });

    // We will also create a module for our other entry point, 'main.zig'.
    const exe_mod = b.createModule(.{
        // `root_source_file` is the Zig "entry point" of the module. If a module
        // only contains e.g. external object files, you can make this `null`.
        // In this case the main source file is merely a path, however, in more
        // complicated build scripts, this could be a generated file.
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });

    // Modules can depend on one another using the `std.Build.Module.addImport` function.
    // This is what allows Zig source code to use `@import("foo")` where 'foo' is not a
    // file path. In this case, we set up `exe_mod` to import `lib_mod`.
    exe_mod.addImport("foo_lib", lib_mod);

    // Now, we will create a static library based on the module we created above.
    // This creates a `std.Build.Step.Compile`, which is the build step responsible
    // for actually invoking the compiler.
    const lib = b.addLibrary(.{
        .linkage = .static,
        .name = "foo",
        .root_module = lib_mod,
    });

    // This declares intent for the library to be installed into the standard
    // location when the user invokes the "install" step (the default step when
    // running `zig build`).
    b.installArtifact(lib);

    // This creates another `std.Build.Step.Compile`, but this one builds an executable
    // rather than a static library.
    const exe = b.addExecutable(.{
        .name = "foo",
        .root_module = exe_mod,
    });

    // This declares intent for the executable to be installed into the
    // standard location when the user invokes the "install" step (the default
    // step when running `zig build`).
    b.installArtifact(exe);

    // This *creates* a Run step in the build graph, to be executed when another
    // step is evaluated that depends on it. The next line below will establish
    // such a dependency.
    const run_cmd = b.addRunArtifact(exe);

    // By making the run step depend on the install step, it will be run from the
    // installation directory rather than directly from within the cache directory.
    // This is not necessary, however, if the application depends on other installed
    // files, this ensures they will be present and in the expected location.
    run_cmd.step.dependOn(b.getInstallStep());

    // This allows the user to pass arguments to the application in the build
    // command itself, like this: `zig build run -- arg1 arg2 etc`
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    // This creates a build step. It will be visible in the `zig build --help` menu,
    // and can be selected like this: `zig build run`
    // This will evaluate the `run` step rather than the default, which is "install".
    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);

    // Creates a step for unit testing. This only builds the test executable
    // but does not run it.
    const lib_unit_tests = b.addTest(.{
        .root_module = lib_mod,
    });

    const run_lib_unit_tests = b.addRunArtifact(lib_unit_tests);

    const exe_unit_tests = b.addTest(.{
        .root_module = exe_mod,
    });

    const run_exe_unit_tests = b.addRunArtifact(exe_unit_tests);

    // Similar to creating the run step earlier, this exposes a `test` step to
    // the `zig build --help` menu, providing a way for the user to request
    // running the unit tests.
    const test_step = b.step("test", "Run unit tests");
    test_step.dependOn(&run_lib_unit_tests.step);
    test_step.dependOn(&run_exe_unit_tests.step);

    // const hash_backend = b.option([]const HashBackend, "hash-backend", "Hash backend to use");

    const upstream = b.dependency("libgit2", .{});
    const linkage = b.option(std.builtin.LinkMode, "linkage", "Link mode") orelse .static;
    const strip = b.option(bool, "strip", "Omit debug information");
    const pic = b.option(bool, "pie", "Produce Position Independent Code");

    const libgit2_util = b.addLibrary(.{
        .linkage = linkage,
        .name = "libgit2_util_lib",
        .root_module = b.createModule(.{
            .target = target,
            .optimize = optimize,
            .pic = pic,
            .strip = strip,
            .link_libc = true,
        }),
    });
    b.installArtifact(libgit2_util);
    libgit2_util.root_module.addIncludePath(upstream.path("include"));
    libgit2_util.root_module.addIncludePath(upstream.path("src/util"));
    libgit2_util.root_module.addCSourceFiles(.{
        .root = upstream.path("src"),
        .files = &.{"alloc.c"},
    });
}

pub const HashBackend = enum {
    builtin,
    openssl,
    openssl_dynamic,
    openssl_fips,
    commoncrypto,
    mbedtls,
    win32,

    pub fn sources_sha1(self: HashBackend) []const u8 {
        switch (self) {
            .builtin => return &.{ "hash/sha1dc/sha1.c", "hash/sha1dc/ubc_check.c", "hash/collisiondetect.c" },
            .openssl => return &.{"hash/openssl.c"},
            .openssl_dynamic => return &.{"hash/openssl.c"},
            .openssl_fips => return &.{"hash/openssl.c"},
            .commoncrypto => return &.{"hash/common_crypto.c"},
            .mbedtls => return &.{"hash/mbedtls.c"},
            .win32 => return &.{"hash/win32.c"},
        }
    }

    pub fn sources_sha256(self: HashBackend) []const u8 {
        switch (self) {
            .builtin => return &.{ "hash/builtin.c", "hash/rfc6234/sha224-256.c" },
            .openssl => return &.{"hash/openssl.c"},
            .openssl_dynamic => return &.{"hash/openssl.c"},
            .openssl_fips => return &.{"hash/openssl.c"},
            .commoncrypto => return &.{"hash/common_crypto.c"},
            .mbedtls => return &.{"hash/mbedtls.c"},
            .win32 => return &.{"hash/win32.c"},
        }
    }
};

const util_sources: []const []const u8 = &.{
    "runtime.c",
    "zstream.c",
    "varint.c",
    "tsort.c",
    "vector.c",
    "pqueue.c",
    "net.c",
    "alloc.c",
    "strlist.c",
    "str.c",
    "wildmatch.c",
    "date.c",
    "regexp.c",
    "sortedcache.c",
    "thread.c",
    "filebuf.c",
    "futils.c",
    "pool.c",
    "fs_path.c",
    "utf8.c",
    "posix.c",
    "hash.c",
    "errors.c",
    "util.c",
    "rand.c",
    "allocators/debugalloc.c",
    "allocators/failalloc.c",
    "allocators/stdalloc.c",
    "allocators/win32_leakcheck.c",
};

const win32_sources: []const []const u8 = &.{
    "win32/w32_util.c",
    "win32/dir.c",
    "win32/map.c",
    "win32/w32_buffer.c",
    "win32/process.c",
    "win32/path_w32.c",
    "win32/error.c",
    "win32/precompiled.c",
    "win32/posix_w32.c",
    "win32/thread.c",
    "win32/w32_leakcheck.c",
    "win32/utf-conv.c",
};

const unix_sources: []const []const u8 = &.{ "unix/map.c", "unix/process.c", "unix/realpath.c" };
