# Author: Lorenzo
# Email : zetatez@icloud.com

echo "uninstall lazy"
rm -f ~/.config/lazy/lazy.yaml
cargo uninstall --bin lazy
echo "done"
