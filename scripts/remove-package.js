const fs = require('fs');
const path = require('path');

const packageName = process.argv[2];
if (!packageName) {
  console.error('❌ Please provide an package name: npm run remove:package my-app');
  process.exit(1);
}

const packagePath = path.join(__dirname, '../packages', packageName);

// Check if the package exists
if (!fs.existsSync(packagePath)) {
  console.error(`❌ Package "${packageName}" does not exist.`);
  process.exit(1);
}

// Delete the package directory
fs.rmSync(packagePath, { recursive: true, force: true });
console.log(`✅ Package "${packageName}" removed successfully.`);
