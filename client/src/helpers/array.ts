// Recursively flattens array.
export function flattenDeep<T>(arr: T): T[] {
  return Array.isArray(arr)
    ? arr.reduce((a, b) => a.concat(flattenDeep(b)), [])
    : [arr];
}
