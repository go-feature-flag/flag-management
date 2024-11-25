import i18n from "i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import * as _ from "lodash";
import { initReactI18next } from "react-i18next";
import { alertComponentTranslation } from "./components/alert/translation.ts";
import { modalComponentTranslation } from "./components/modal/translation.ts";
import { navigationComponentTranslation } from "./components/navigation/translation.ts";
import { contactPageTranslation } from "./routes/authenticatedRoutes/contact/translation.ts";
import { errorPageTranslation } from "./routes/authenticatedRoutes/error/translation.ts";
import { flagPageTranslation } from "./routes/authenticatedRoutes/flag/translation.ts";
import { flagsPageTranslation } from "./routes/authenticatedRoutes/flags/translation.ts";
i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    debug: true,
    fallbackLng: "en",
    interpolation: {
      escapeValue: false,
    },
    resources: _.merge(
      contactPageTranslation,
      flagsPageTranslation,
      flagPageTranslation,
      errorPageTranslation,
      alertComponentTranslation,
      modalComponentTranslation,
      navigationComponentTranslation,
    ),
  });
export default i18n;
