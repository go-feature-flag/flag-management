export interface colorSelector {
  emoji: string;
  color: string;
}

export function getColorByIndex(index: number): colorSelector {
  const emojiColors: colorSelector[] = [
    { emoji: "🟢", color: "bg-green-500" },
    { emoji: "🔵", color: "bg-blue-500" },
    { emoji: "🟡", color: "bg-yellow-300" },
    { emoji: "🟤", color: "bg-amber-800" },
    { emoji: "🔴", color: "bg-red-600" },
    { emoji: "⚪", color: "bg-slate-50" },
    { emoji: "⚫", color: "bg-gray-950" },
    { emoji: "🟠", color: "bg-orange-500" },
    { emoji: "🟣", color: "bg-purple-800" },
  ];

  const positiveIndex = index < 0 ? index + emojiColors.length : index;
  const colorIndex = positiveIndex % emojiColors.length;
  return emojiColors[colorIndex];
}
