import { useBooleanFlagValue } from "@openfeature/react-sdk";
import { Alert, Tabs } from "flowbite-react";
import { Suspense, useState } from "react";
import type {
  FieldErrors,
  SubmitErrorHandler,
  SubmitHandler,
} from "react-hook-form";
import { FormProvider, useForm, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { BiCodeBlock } from "react-icons/bi";
import { FaInfoCircle } from "react-icons/fa";
import { GrDeploy } from "react-icons/gr";
import { HiAdjustments } from "react-icons/hi";
import { IoSettingsOutline } from "react-icons/io5";
import { MdOutlineDataObject } from "react-icons/md";
import { TbTargetArrow } from "react-icons/tb";
import {
  Await,
  useBlocker,
  useLoaderData,
  useLocation,
  useNavigate,
  useParams,
} from "react-router-dom";
import { v4 as uuidv4 } from "uuid";
import { AlertError } from "../../../components/alert/error.tsx";
import Loader from "../../../components/loader/loader.tsx";
import { config } from "../../../config.ts";
import { useDocumentTitle } from "../../../hooks/documentTitle.ts";
import type { FeatureFlagFormData } from "../../../models/featureFlagFormData.ts";
import {
  createFlag,
  updateFlagById,
} from "../../../service/featureFlagService.ts";
import { FlagLayout } from "./FlagLayout.tsx";
import { CodeExamples } from "./code-examples/codeExampleTab.tsx";
import { variationTypes } from "./helpers/variations.ts";
import { Metadata } from "./metadata/metadata.tsx";
import { ExitWithoutSaveModal } from "./modals/exitWithoutSaveModal.tsx";
import { FlagSaveButton } from "./saveButton.tsx";
import { Settings } from "./settings/settings.tsx";
import { TargetingPage } from "./targeting/targetingPage.tsx";
import { VariationTab } from "./variations/variationTab.tsx";

const translationBaseKey = "page.flag";

/**
 * FlagPage is the wrapper for the flag page and loads the data from the API.
 */
export const FlagPage = ({ isNew = false }: { isNew?: boolean }) => {
  const { t } = useTranslation();
  useDocumentTitle(
    isNew
      ? t(`${translationBaseKey}.pageTitleNew`)
      : t(`${translationBaseKey}.pageTitleEdit`),
  );
  const data = useLoaderData() as { result: FeatureFlagFormData };

  return (
    <Suspense fallback={<FlagPageLoader />}>
      <Await resolve={data.result} errorElement={<ErrorLoadingPage />}>
        {(featureFlags) => (
          <FlagPageContent flag={featureFlags} isNew={isNew} />
        )}
      </Await>
    </Suspense>
  );
};

const FlagPageContent = ({
  flag,
  isNew = false,
}: {
  flag: FeatureFlagFormData;
  isNew?: boolean;
}) => {
  const [displayToastError, setDisplayToastError] = useState(false);
  const [displayToastSuccess, setDisplayToastSuccess] = useState(false);
  const [loadingSave, setLoadingSave] = useState(false);
  const navigate = useNavigate();
  const [hasError, setHasError] = useState<boolean>(false);
  const [mainErrors, setMainErrors] = useState<
    FieldErrors<FeatureFlagFormData>
  >({});

  // Init name of the flag if we are in the creation of a new flag
  const urlParams = useParams<{ flagName: string }>();
  if (isNew) {
    flag.name = urlParams.flagName;
    flag.id = uuidv4();
  }

  const methods = useForm<FeatureFlagFormData>({
    defaultValues: flag,
  });

  const onError: SubmitErrorHandler<FeatureFlagFormData> = (errors) => {
    setDisplayToastError(true);
    setTimeout(() => {
      setDisplayToastError(false);
    }, 3000);
    setHasError(true);
    setMainErrors(errors);
    console.log("onError", errors);
  };
  const onSubmit: SubmitHandler<FeatureFlagFormData> = async (
    data: FeatureFlagFormData,
  ) => {
    setHasError(false);
    setLoadingSave(true);

    if (isNew) {
      const response = await createFlag(data);
      methods.reset(data, { keepValues: true });
      navigate(`/flags/${response.id}`);
    } else {
      await updateFlagById(data);
      methods.reset(data, { keepValues: true });
    }
    setLoadingSave(false);
    setDisplayToastSuccess(true);
    setTimeout(() => {
      setDisplayToastSuccess(false);
    }, config.flagPageConfiguration.successToastDisplayTime);
  };

  return (
    <>
      <FlagLayout name={flag.name ?? ""}>
        <FormProvider {...methods}>
          <form onSubmit={methods.handleSubmit(onSubmit, onError)}>
            <NavigationConfirmationModal isDirty={methods.formState.isDirty} />
            <FlagSaveButton
              isDirty={methods.formState.isDirty}
              displayToastError={displayToastError}
              setDisplayToastError={setDisplayToastError}
              displayToastSuccess={displayToastSuccess}
              setDisplayToastSuccess={setDisplayToastSuccess}
              loadingSave={loadingSave}
            />
            <FeatureFlagTabs
              isNew={isNew}
              hasError={hasError}
              mainErrors={mainErrors}
            />
          </form>
        </FormProvider>
      </FlagLayout>
    </>
  );
};

const ErrorMessage = ({
  mainErrors,
}: {
  mainErrors: FieldErrors<FeatureFlagFormData>;
}) => {
  const renderErrors = (errors: FieldErrors, parentKey = ""): JSX.Element[] => {
    console.log("errors", errors);

    return Object.entries(errors).flatMap(([key, value]) => {
      const fullKey = parentKey ? `${parentKey}.${key}` : key;
      if (value?.message) {
        return <li key={fullKey}>{value.message as string}</li>;
      } else if (typeof value === "object") {
        return renderErrors(value as FieldErrors, fullKey);
      } else {
        return [];
      }
    });
  };

  return (
    <Alert color="failure">
      <div className={"flex space-x-2"}>
        <FaInfoCircle className={"h-4 w-4"} />
        <span className="font-medium">Error in your flag configuration !</span>
      </div>
      <ul className="ml-5 mt-1.5 list-inside list-disc">
        {renderErrors(mainErrors)}
      </ul>
    </Alert>
  );
};
/**
 * FlagPageLoader is the loading state of the flag page.
 */
const FlagPageLoader = () => {
  return (
    <FlagLayout name={"Loading ..."}>
      <Loader />
    </FlagLayout>
  );
};

const NavigationConfirmationModal = ({ isDirty }: { isDirty: boolean }) => {
  const blocker = useBlocker(({ currentLocation, nextLocation }) => {
    if (
      currentLocation.pathname.startsWith("/flags/new") &&
      nextLocation.pathname.startsWith("/flags/")
    ) {
      return false;
    }
    return isDirty && currentLocation.pathname !== nextLocation.pathname;
  });
  return (
    <>
      {blocker.state === "blocked" ? (
        <ExitWithoutSaveModal
          isOpen={true}
          onClickCancel={() => blocker.reset()}
          onClickYes={() => blocker.proceed()}
        />
      ) : null}
    </>
  );
};
/**
 * ErrorLoadingPage is the error state of the flag page.
 */
const ErrorLoadingPage = () => {
  const { t } = useTranslation();
  return (
    <FlagLayout name={"Error"}>
      <AlertError
        text={t(`${translationBaseKey}.error.dataFetch`)}
        proposeRefresh={true}
      />
    </FlagLayout>
  );
};

/**
 * FeatureFlagTabs is the tabs section for the flag page.
 */
const FeatureFlagTabs = ({
  isNew,
  hasError,
  mainErrors,
}: {
  isNew: boolean;
  hasError: boolean;
  mainErrors: FieldErrors<FeatureFlagFormData>;
}) => {
  const location = useLocation();
  const { type } = useWatch<FeatureFlagFormData>();
  const getFlagType = (): string => {
    if (!isNew) {
      return type as string;
    }

    const queryParams = new URLSearchParams(location.search);
    return queryParams.get("type") ?? variationTypes[0].type;
  };

  const enableRolloutTab = useBooleanFlagValue("app-enable-rollout-tab", false);
  return (
    <>
      <Tabs id={"flag_type"}>
        <Tabs.Item active title="Variations" icon={HiAdjustments}>
          {hasError && <ErrorMessage mainErrors={mainErrors} />}
          <VariationTab type={getFlagType()} />
        </Tabs.Item>
        <Tabs.Item title="Targeting" icon={TbTargetArrow}>
          {hasError && <ErrorMessage mainErrors={mainErrors} />}
          <TargetingPage />
        </Tabs.Item>
        {enableRolloutTab && (
          <>
            {hasError && <ErrorMessage mainErrors={mainErrors} />}
            <Tabs.Item title="Rollout" icon={GrDeploy}></Tabs.Item>
          </>
        )}
        <Tabs.Item title="Metadata" icon={MdOutlineDataObject}>
          {hasError && <ErrorMessage mainErrors={mainErrors} />}
          <Metadata />
        </Tabs.Item>
        <Tabs.Item title={"Settings"} icon={IoSettingsOutline}>
          {hasError && <ErrorMessage mainErrors={mainErrors} />}
          <Settings />
        </Tabs.Item>
        <Tabs.Item title="Code Example" icon={BiCodeBlock}>
          <CodeExamples />
        </Tabs.Item>
      </Tabs>
    </>
  );
};
