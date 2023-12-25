export default function isNonEmptyString(value: string | null | undefined): value is string {
  if (value == null) {
    return false;
  }
  return value.length > 0;
}
