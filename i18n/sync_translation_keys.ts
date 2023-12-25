import * as fs from 'fs/promises';
import * as path from 'path';

type Translations = Record<string, string>;

async function syncTranslationKeys(
  sourceFilePath: string, 
  targetFilePath: string,
): Promise<void> {
  const sourceJson: Translations = JSON.parse(await fs.readFile(sourceFilePath, 'utf-8'));
  const targetJson: Translations = JSON.parse(await fs.readFile(targetFilePath, 'utf-8'));

  const newJson: Translations = {};

  let count = 0
  for (const key of Object.keys(sourceJson)) {
    const existing = targetJson[key];
    if (existing == null) {
      ++count;
    }

    newJson[key] =  existing ?? `<UNTRANSLATED> ${sourceJson[key]}`;
  }

  console.log(`Synced ${count} new keys to ${path.basename(targetFilePath)}`)
  await fs.writeFile(targetFilePath, JSON.stringify(newJson, null, 2));
}

async function getTranslationFiles(dirPath: string): Promise<Array<string>> {
  const files = await fs.readdir(dirPath);
  const jsonFiles = files.filter(file => path.extname(file) == '.json');
  return jsonFiles.map(file => path.join(dirPath, file));
}

async function main(): Promise<void> {
  const dir = path.join(__dirname, 'translations');
  const sourceFilePath = path.join(__dirname, 'translations', 'en-US.json');
  const translationFilePaths = (await getTranslationFiles(dir)).filter(file => file !== sourceFilePath);
  await Promise.all(translationFilePaths.map(translationPath => syncTranslationKeys(sourceFilePath, translationPath)));
}

main();