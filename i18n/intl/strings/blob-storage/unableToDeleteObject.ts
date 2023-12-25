import intl from '../../intl';

export default function unableToDeleteObject(): string {
  return intl.get('blob-storage.unable-to-delete-object').d('Unable to delete');
}