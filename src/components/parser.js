const fs = require('fs');
const path = require('path');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
    this.content = '';
  }

  async readContent() {
    try {
      this.content = await fs.promises.readFile(this.filePath, 'utf8');
    } catch (error) {
      throw new Error(`Failed to read file: ${error.message}`);
    }
  }

  parseContent() {
    if (!this.content) {
      throw new Error('No content to parse');
    }
    const lines = this.content.split('\n');
    const parsedData = lines.map(line => line.trim());
    return parsedData;
  }

  async parseFile() {
    await this.readContent();
    return this.parseContent();
  }
}

module.exports = Parser;