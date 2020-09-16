module.exports = {
  root: true,
  parser: "@typescript-eslint/parser",
  parserOptions: {},
  plugins: ["@typescript-eslint"],
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "prettier/@typescript-eslint",
    "plugin:eslint-comments/recommended",
  ],
  env: {
    node: true,
  },
  rules: {
    "eslint-comments/no-unused-disable": "error",
    "@typescript-eslint/explicit-module-boundary-types": "off",
  },
};
