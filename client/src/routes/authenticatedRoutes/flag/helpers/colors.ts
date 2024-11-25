export function getColorByIndex(index: number): string {
  const emojiColors: string[] = [
    "🟢",
    "🔵",
    "🟡",
    "🟤",
    "🔴",
    "⚪",
    "⚫",
    "🟠",
    "🟣",
    "⭕",
  ];

  const positiveIndex = index < 0 ? index + emojiColors.length : index;
  const colorIndex = positiveIndex % emojiColors.length;
  return emojiColors[colorIndex];
}
