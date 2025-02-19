package main

import (
	"context"
	"dagger/go-fiber/internal/dagger"
)

type GoFiber struct{}

// Returns a container that echoes whatever string argument is provided
func (m *GoFiber) ContainerEcho(ctx context.Context, stringArg string) (string, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg}).Stdout(ctx)
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *GoFiber) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

// Run application unit tests
func (m *GoFiber) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.BaseEnv(source).
		WithExec([]string{"go", "test", "-v"}).
		Stdout(ctx)
}

// Sets up the base Go environment
func (m *GoFiber) BaseEnv(source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.23").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		// WithExec([]string{"go", "mod", "download"}).
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-123"))
}

// ... existing code ...

// func (m *GoFiber) Test(ctx context.Context, source *dagger.Directory) (string, error) {
// 	return dag.Go().
// 		Container().
// 		WithMountedDirectory("/src", source).
// 		WithWorkdir("/src").
// 		WithExec([]string{"go", "mod", "download"}).
// 		WithExec([]string{"go", "test", "-v"}).
// 		Stdout(ctx)
// }
