import {readFile} from 'fs/promises';

export default async function loadSchema(): Promise<string> {
  return await readFile('./schema.graphqls', 'utf-8');
}