import intl from '../../intl';

export default function unableToCreateDownloadUrl(): string {
  return intl.get('blob-storage.unable-to-create-download-url').d('Unable to create download url');
}