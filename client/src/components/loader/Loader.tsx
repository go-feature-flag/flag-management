import { Spinner } from "flowbite-react";

export interface LoaderProps {
  alt?: string;
  className?: string;
}

export default function Loader(
  { alt, className }: LoaderProps = { alt: "Loading..." },
) {
  return <Spinner aria-label={alt} className={className} />;
}
