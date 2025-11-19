const fs = require('fs');
const path = require('path');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
    this.fileContent = '';
  }

  async readFile() {
    try {
      this.fileContent = await fs.promises.readFile(this.filePath, 'utf8');
    } catch (error) {
      throw new Error(`Failed to read file: ${error.message}`);
    }
  }

  parseFileContent() {
    if (!this.fileContent) {
      throw new Error('File content is empty');
    }

    const lines = this.fileContent.split('\n');
    const parsedData = [];

    lines.forEach((line) => {
      const trimmedLine = line.trim();
      if (trimmedLine) {
        const [key, value] = trimmedLine.split('=');
        if (key && value) {
          parsedData.push({ key: key.trim(), value: value.trim() });
        }
      }
    });

    return parsedData;
  }
}

module.exports = Parser;