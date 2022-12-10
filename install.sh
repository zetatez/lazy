# Author: Lorenzo
# Email : zetatez@icloud.com

echo "install lazy"
rm -rf ~/.config/lazy && mkdir -p ~/.config/lazy
cp -f lazy.yaml ~/.config/lazy/
cargo install --path .
# sudo cp ./target/release/lazy /usr/local/bin/
echo "done"
