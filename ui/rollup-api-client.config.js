import amd from 'rollup-plugin-amd';

export default {
  input: '../api/js/client.js',
  output: {
    file: 'src/api/client.js',
    format: 'esm'
  },
  plugins: [
    amd()
  ]
};
