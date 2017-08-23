package templates

// Script to lint
const Script = `
let Linter = require('eslint').Linter;
let chalk = require('chalk')
let linter = new Linter();

let logErrors = function(err) {
	let snippet = chalk.yellow('	line no:' + err.line)  + ' ' + err.message;
	console.log(snippet);
}

let errors = linter.verify('{{.Code}}' ,{{.LintConfig}},  { filename: '{{.FileName}}' });
if (errors.length) {
	console.log(chalk.cyan('Errors in file {{.FileName}} :-' ))
	errors.forEach( err => logErrors(err))
}
`
