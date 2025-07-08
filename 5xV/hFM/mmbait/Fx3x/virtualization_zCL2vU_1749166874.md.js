const env = process.env.NODE_ENV;

  module.exports = {
    ignore: [
      '*.jest.js',
      '*.e2e.js',
      '*.ssr.js',
      '*.example.js',
      'source/demo',
      'source/jest-*.js',
      'source/TestUtils.js',
      '@babel/plugin-transform-runtime',
      '@babel/plugin-proposal-class-properties',
      '@babel/plugin-transform-flow-comments',
      ['flow-react-proptypes', {deadCode: true, useESModules: true}],
      ['transform-react-remove-prop-types', {mode: 'wrap'}],
    ],
    presets: [
      ['@babel/preset-env', {modules: false}],
      '@babel/preset-react',

    module.exports.plugins.push('@babel/plugin-transform-modules-commonjs');
  }
}

if (env === 'rollup') {
  module.exports = {
    plugins: [
      '@babel/plugin-external-helpers',
      '@babel/plugin-proposal-class-properties',
    ],
    presets: [
      ['@babel/preset-env', {modules: false}],
      '@babel/preset-react',
      '@babel/preset-flow',
  };

if (env === 'development') {
  module.exports = {
    plugins: ['@babel/plugin-proposal-class-properties'],
    presets: ['@babel/preset-react', '@babel/preset-flow'],
  };
}
  module.exports = {
    comments: false,
    plugins: [
      '@babel/plugin-transform-runtime',
      '@babel/plugin-proposal-class-properties',
    ],
    presets: ['@babel/preset-env', '@babel/preset-react', '@babel/preset-flow'],
}

  module.exports = {
    comments: false,

    plugins: [
      '@babel/plugin-transform-modules-commonjs',
      '@babel/plugin-proposal-class-properties',
    ],
    presets: ['@babel/preset-react', '@babel/preset-flow'],
}
