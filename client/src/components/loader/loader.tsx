import { Spinner } from "flowbite-react";
import { t } from "i18next";

export interface LoaderProps {
  alt?: string;
  className?: string;
}

const translationBaseKey = "component.loader";
export default function Loader(
  { alt, className }: LoaderProps = { alt: t(`${translationBaseKey}.loading`) },
) {
  return <Spinner aria-label={alt} className={className} />;
}
