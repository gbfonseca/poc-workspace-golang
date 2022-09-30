const {build} = require('esbuild');
const fs = require('fs-extra');

const generateBuild = async () => {

    await build({
        entryPoints: ['./src/index.js'],
        outdir: './build',
        minify: true,
        bundle: true,
        sourcemap: true,
        target:['es6', 'es2016'],
        loader: {
            ".js": "jsx",
            ".svg": "file",
            ".png": "file",
            ".jpg": "file",
            ".jpeg": "file",
            ".gif": "file"
        },
        define: {
            'process.env.NODE_ENV': "'development'",
        }
    }).catch(() => process.exit(1));

};
generateBuild();
