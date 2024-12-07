import { useEffect } from "react";
import { config } from "../config.ts";
export const useDocumentTitle = (title: string) => {
  useEffect(() => {
    document.title = `${title} | ${config.title}`;
  }, [title]);
  return null;
};
