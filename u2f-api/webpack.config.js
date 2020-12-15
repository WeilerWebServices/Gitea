const path = require('path');

module.exports = {
    entry: './dist/lib/u2f-api.js',
    output: {
        path: path.resolve(__dirname, 'web'),
        filename: 'u2f.bundle.js',
        library: "u2fApi",
        libraryTarget: "this",
    },
    mode: "production"
};
