# Get the version.
version=`git describe --tags`
# Write out the package.
cat << EOF > version.go
package version
//go:generate sh get_version.sh
var Version = "$version"
EOF