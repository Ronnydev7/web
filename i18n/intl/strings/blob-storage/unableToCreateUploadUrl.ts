import intl from '../../intl';

export default function unableToCreateUploadUrl(): string {
  return intl.get('blob-storage.unable-to-create-upload-url').d('Unable to generate upload URL');
}