// node-resolve will resolve all the node dependencies
import resolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs'
import babel from 'rollup-plugin-babel'
import replace from 'rollup-plugin-replace'
import serve from 'rollup-plugin-serve'

export default {
  input: 'src/index.js',
  output: {
    file: 'dist/bundle.js',
    format: 'iife'
  },
  // All the used libs needs to be here
  // external: [
  //   'react',
  //   'react-proptypes'
  // ],
  plugins: [
    babel({
      exclude: 'node_modules/**'
    }),
    commonjs({
      exclude: 'node_modules/process-es6/**',
      include: [
        //some react related modules i need
        'node_modules/fbjs/**',
        'node_modules/object-assign/**',
        'node_modules/react/**',
        'node_modules/react-dom/**',
        'node_modules/prop-types/**',
        'node_modules/create-react-class/**' // adding the module with that "default not exported by" message to this includes list, made that message go away
      ]
    }),
    resolve(),
    replace({
      'process.env.NODE_ENV': JSON.stringify( 'production' )
    }),
    // https://github.com/thgh/rollup-plugin-serve#options
    serve({
      contentBase: 'dist',
      port: 8080
    })
  ]
};
