const {exec, execSync} = require('child_process');
const fs = require('fs');
const assert = require('assert');

const output = "toto.stl.xml"
const expected = "./sources_100/1.stl.xml"
const src = "./sources_100/1.stl"

const {format} = require('prettier')
//exec(`go run . --input ${input} --output ${output}`, (error, stdout, stderr) => {

describe('', () => {
  it('should toto', () => {
    execSync(`go run . --input ${src} --output ${output}`);
    const file1 = format(fs.readFileSync(expected, 'utf8'), {
      parser: 'html',
      printWidth: 1000,
      htmlWhitespaceSensitivity: 'ignore'
    })
    const file2 = format(fs.readFileSync(output, 'utf8'), {
      parser: 'html',
      printWidth: 1000,
      htmlWhitespaceSensitivity: 'ignore'
    })

// Compare the two files and show the difference
    expect(file1).toEqual(file2)
  });
});