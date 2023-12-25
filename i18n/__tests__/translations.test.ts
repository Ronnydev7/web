import enUS from '../translations/en-US.json';
import zhCN from '../translations/zh-CN.json';

describe('translations', () => {
  const sourceOfTruth = enUS;
  const sourceOfTruthKeysSorted = Object.keys(sourceOfTruth).sort();
  
  describe('translation keys match the source of truth', () => {
    type TestCase = Readonly<{
      translation: typeof enUS,
    }>;
    const testCases: Array<[string, TestCase]> = [
      [
        'zhCN',
        {
          translation: zhCN,
        },
      ],
    ];

    it.each(testCases)(
      '%s',
      (_, {translation}) => {
        expect(Object.keys(translation).sort()).toStrictEqual(sourceOfTruthKeysSorted);
      },
    );
  });
});