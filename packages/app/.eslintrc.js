module.exports = {
  root: true,
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
  },
  plugins: ["@typescript-eslint", "jest", "react-hooks", "react"],
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:jest/recommended",
    "prettier/@typescript-eslint",
    "plugin:eslint-comments/recommended",
    "plugin:react-hooks/recommended",
    "plugin:react/recommended",
  ],
  env: {
    node: true,
    jest: true,
  },
  rules: {
    "eslint-comments/no-unused-disable": "error",
    "react-hooks/rules-of-hooks": "error",
    "react-hooks/exhaustive-deps": "warn",
  },
};
