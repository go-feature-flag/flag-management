import type { FC } from "react";

export interface ImageProps {
  src: string;
  alt: string;
  height?: number;
  width?: number;
  className?: string;
}
const Image: FC<ImageProps> = (props: ImageProps) => {
  return (
    <img
      src={props.src}
      alt={props.alt}
      height={props.height}
      width={props.width}
      className={props.className}
    />
  );
};

export default Image;
