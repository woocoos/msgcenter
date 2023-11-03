const { getESLintConfig } = require('@applint/spec');

module.exports = getESLintConfig(
  'react-ts', {
  rules: {
    "react-hooks/exhaustive-deps": 0
  }
}
);
