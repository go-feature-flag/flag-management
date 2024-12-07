import { useTranslation } from "react-i18next";
import { BsToggleOn } from "react-icons/bs";
import { useDocumentTitle } from "../../../hooks/documentTitle.ts";
import { FlagList } from "./flagList.tsx";

const translationBaseKey = "page.flags";
export const FlagListPage = () => {
  useDocumentTitle("Feature Flags");
  const { t } = useTranslation();
  return (
    <div className="grid place-items-center p-3 md:p-6 lg:p-10">
      <div className="w-full max-w-7xl grid-cols-2">
        <div className={"flex flex-row"}>
          <BsToggleOn className={"mr-3 text-4xl"} />
          <h1>{t(`${translationBaseKey}.title`)}</h1>
        </div>
        <p className="pageDescription">
          {t(`${translationBaseKey}.subtitle.line1`)} <br />
          {t(`${translationBaseKey}.subtitle.line2`)}
        </p>
        <hr />
        <FlagList />
      </div>
    </div>
  );
};
