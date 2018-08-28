// node-resolve will resolve all the node dependencies
// import resolve from 'rollup-plugin-node-resolve';
import babel from 'rollup-plugin-babel';

export default {
  input: 'src/index.js',
  output: {
    file: 'dist/bundle.js',
    format: 'cjs'
  },
  external: [
    'react'
  ],
  plugins: [
    babel({
      exclude: 'node_modules/**'
    })
  ]
};
