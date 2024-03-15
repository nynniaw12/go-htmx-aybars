/** @type {import('tailwindcss').Config} */
const plugin = require('tailwindcss/plugin')

module.exports = {
    content: ["./templates/**/*.{html,js,templ,go}",
        "./main.go"],
    theme: {
        extend: {
            transitionDuration: {
                '5000': '5000ms',
            },
            spacing: {
                '0': '0px',
                '1': '0.25rem',
                '2': '0.5rem',
                '3': '0.75rem',
                '4': '1rem', 
            }
        },
        fontFamily: {
            serif: ["EditorialNew"],
            sans: ["Quicksand"]
        },
    },
    plugins: [
        require('tailwindcss'),
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
        plugin(function({ addVariant }) {
            addVariant('htmx-settling', ['&.htmx-settling', '.htmx-settling &'])
            addVariant('htmx-request', ['&.htmx-request', '.htmx-request &'])
            addVariant('htmx-swapping', ['&.htmx-swapping', '.htmx-swapping &'])
            addVariant('htmx-added', ['&.htmx-added', '.htmx-added &'])
        }),
    ],
}
