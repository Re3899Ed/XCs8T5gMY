/**
 * CRACO configuration JS file
 */

module.exports = {
        postcss: {
            plugins: [
                require('tailwindcss'),
                require('autoprefixer'),
            ],
        },
    },
}