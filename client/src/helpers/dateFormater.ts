export function formatDate(inputDate?: Date): string {
  if (!inputDate) {
    return "";
  }

  const options: Intl.DateTimeFormatOptions = {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  };

  return new Intl.DateTimeFormat("en-US", options).format(inputDate);
}
