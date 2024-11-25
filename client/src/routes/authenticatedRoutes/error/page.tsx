import { useTranslation } from "react-i18next";
import type { ErrorResponse } from "react-router-dom";
import { Link, isRouteErrorResponse, useRouteError } from "react-router-dom";
import img404 from "../../../assets/pages/error/404.svg";
import img500 from "../../../assets/pages/error/500.svg";
import Image from "../../../components/image/Image";
import { useDocumentTitle } from "../../../hooks/documentTitle.ts";

const baseTranslationKey = "page.error";
export const ErrorHandler = () => {
  const { t } = useTranslation();
  useDocumentTitle(t(`${baseTranslationKey}.pageTitle`));
  const error = useRouteError() as ErrorResponse;
  if (isRouteErrorResponse(error)) {
    if (error.status === 404) {
      return (
        <ErrorPage
          title={"404"}
          detail={t(`${baseTranslationKey}.err404`)}
          img={img404}
          imgAlt={"Page not found"}
        />
      );
    }
  }

  return (
    <ErrorPage
      title={"Error"}
      detail={t(`${baseTranslationKey}.errDefault`)}
      img={img500}
      imgAlt={"Error"}
    />
  );
};

interface ErrorProps {
  title: string;
  detail: string;
  img: string;
  imgAlt: string;
  errorMessage?: string;
}

const ErrorPage = (props: ErrorProps) => {
  const { t } = useTranslation();
  return (
    <div className="flex h-fit w-fit items-center">
      <div className="container flex flex-col items-center justify-between px-5 md:flex-row">
        <div className="mx-8 w-full lg:w-1/2">
          <div className={"mb-8 text-7xl font-extrabold text-goff-500"}>
            {props.title}
          </div>
          <p className="mb-8 text-2xl leading-normal md:text-3xl">
            {props.detail}
          </p>
          <p>
            {props.errorMessage && (
              <span className="text-sm text-red-500">{props.errorMessage}</span>
            )}
          </p>
          <Link
            to="/"
            className="inline rounded-lg border border-transparent bg-goff-600 px-5 py-3 text-sm font-medium leading-5 text-white shadow-2xl transition-all hover:bg-goff-700 focus:outline-none active:bg-goff-800"
          >
            {t(`${baseTranslationKey}.backButton`)}
          </Link>
        </div>
        <div className="mx-5 my-12 w-full lg:flex lg:w-1/2 lg:justify-end">
          <Image src={props.img} alt={props.imgAlt} width={450} />
        </div>
      </div>
    </div>
  );
};
