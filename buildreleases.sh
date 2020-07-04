#!/bin/zsh
# run this from the root project folder (canaveral)

# build and move new binaries into release folders
echo "Building canaveral_darwin_amd64..."
env GOOS=darwin GOARCH=amd64 go build
mv canaveral dist/canaveral_darwin_amd64/
echo "done.\n"

echo "Building canaveral_linux_386..."
env GOOS=linux GOARCH=386 go build
mv canaveral dist/canaveral_linux_386/
echo "done.\n"

echo "Building canaveral_linux_amd64..."
env GOOS=linux GOARCH=amd64 go build
mv canaveral dist/canaveral_linux_amd64/
echo "done.\n"

echo "Building canaveral_windows_386..."
env GOOS=windows GOARCH=386 go build
mv canaveral.exe dist/canaveral_windows_386/
echo "done.\n"

echo "Building canaveral_windows_amd64..."
env GOOS=windows GOARCH=amd64 go build
mv canaveral.exe dist/canaveral_windows_amd64/
echo "done.\n"

# Create new tarballs
cd dist
echo "Creating tarballs..."
tar -czf canaveral_darwin_amd64.tar.gz canaveral_darwin_amd64
tar -czf canaveral_linux_386.tar.gz canaveral_linux_386
tar -czf canaveral_linux_amd64.tar.gz canaveral_linux_amd64
tar -czf canaveral_windows_386.tar.gz canaveral_windows_386
tar -czf canaveral_windows_amd64.tar.gz canaveral_windows_amd64
echo "done"
