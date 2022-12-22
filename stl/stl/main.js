const {exec} = require('child_process');
const fs = require('fs');
const assert = require('assert');

const input = "toto.xml"
const output = "1.stl.xml"

const {format} = require('prettier')
//exec(`go run . --input ${input} --output ${output}`, (error, stdout, stderr) => {


// Read the two files
const file1 = format(fs.readFileSync(input, 'utf8'), {parser: 'html'})
const file2 = format(fs.readFileSync(output, 'utf8'), {parser: 'html'})

// Compare the two files and show the difference
// try {
//   assert.strictEqual(file1, file2);
//   console.log('The files are the same');
// } catch (error) {
//   console.log('The files are different:');
//   console.log(error);
// }
//});


describe('', () => {
  it('should toto', () => {
    const file1 = format(fs.readFileSync(input, 'utf8'), {
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