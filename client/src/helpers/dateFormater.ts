export function formatAndLocalizedDate(
  inputDate?: Date,
  locales?: Intl.LocalesArgument,
  options?: Intl.DateTimeFormatOptions,
): string {
  if (!inputDate) {
    return "";
  }
  const formatOptions: Intl.DateTimeFormatOptions = {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  };
  const localDate = new Date(inputDate.toLocaleString(locales, options));
  return new Intl.DateTimeFormat("us-US", formatOptions).format(localDate);
}
