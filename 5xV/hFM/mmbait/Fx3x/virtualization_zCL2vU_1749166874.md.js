  module.exports = {
      '*.jest.js',
      '*.example.js',
      'source/demo',
      'source/TestUtils.js',
      '@babel/plugin-transform-runtime',
      '@babel/plugin-proposal-class-properties',
      '@babel/plugin-transform-flow-comments',
      ['flow-react-proptypes', {deadCode: true, useESModules: true}],
      ['transform-react-remove-prop-types', {mode: 'wrap'}],
      '@babel/preset-react',

  }
}

if (env === 'rollup') {
      '@babel/plugin-proposal-class-properties',
    ],
      ['@babel/preset-env', {modules: false}],
      '@babel/preset-react',
      '@babel/preset-flow',
  };

if (env === 'development') {
    presets: ['@babel/preset-react', '@babel/preset-flow'],
  };
}
    comments: false,
      '@babel/plugin-transform-runtime',
      '@babel/plugin-proposal-class-properties',

  module.exports = {
    comments: false,

    plugins: [
      '@babel/plugin-proposal-class-properties',
    ],
    presets: ['@babel/preset-react', '@babel/preset-flow'],
}
