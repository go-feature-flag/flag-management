import { clsx } from "clsx";
import { Tooltip } from "flowbite-react";
import type { ReactNode } from "react";
import { useTranslation } from "react-i18next";
import { FaCheck } from "react-icons/fa6";
import { VscError } from "react-icons/vsc";
import type { FormFlagPageVariationInfo } from "../../models/featureFlagFormData.ts";
import { getColorByIndex } from "../../routes/authenticatedRoutes/flag/helpers/colors.ts";

const baseTranslationKey = "component.progressBar";
export const PercentageProgressBar = ({
  variations,
  percentages,
}: {
  variations: FormFlagPageVariationInfo[];
  percentages: Record<string, number> | undefined;
}) => {
  return (
    <div className={"flex space-x-2"}>
      <div className={"grow"}>
        <div className="relative min-w-full pt-1">
          <div
            className="mb-4 flex h-4 overflow-hidden rounded bg-gray-300 text-xs dark:bg-gray-600"
            data-testid={"percentage-progress-bar"}
          >
            {variations.map((variation, index) => {
              let percentage = percentages?.[variation.name] ?? 0;
              percentage = Number.isNaN(percentage) ? 0 : percentage;
              return (
                percentage > 0 && (
                  <div
                    key={variation.name}
                    style={{ width: `${percentage}%` }}
                    data-testid={"percentage-progress-bar-item"}
                    className={clsx(
                      "flex flex-col justify-center whitespace-nowrap text-center text-white shadow-none",
                      getColorByIndex(index).color,
                    )}
                  >
                    {percentage > variation.name.length && variation.name}
                  </div>
                )
              );
            })}
          </div>
        </div>
      </div>
      <div className={"w-fit"} data-testid={"percentage-progress-display"}>
        <PercentageValue percentages={percentages} variations={variations} />
      </div>
    </div>
  );
};

const PercentageValue = ({
  percentages,
  variations,
}: {
  percentages: Record<string, number> | undefined;
  variations: FormFlagPageVariationInfo[];
}) => {
  const { t } = useTranslation();

  const sum = variations.reduce((acc, variation) => {
    let value = percentages?.[variation.name] ?? 0;
    value = Number.isNaN(value) ? 0 : value;
    return acc + value;
  }, 0);

  function displayPercentage(sum: number): ReactNode {
    if (Number.isNaN(sum) || sum > 100 || sum < 0) {
      return (
        <Tooltip content={t(`${baseTranslationKey}.errorTooltip`)}>
          <VscError
            className={"mt-0.5 h-5 w-5 text-red-500"}
            data-testid={"percentage-progress-bar-error-icon"}
          />
        </Tooltip>
      );
    }
    if (sum === 100) {
      return (
        <FaCheck
          className={"mt-0.5 h-5 w-5 text-goff-400"}
          data-testid={"percentage-progress-bar-check-icon"}
        />
      );
    }
    return <span className={"text-sm"}>{sum}%</span>;
  }

  return <div>{displayPercentage(sum)}</div>;
};
