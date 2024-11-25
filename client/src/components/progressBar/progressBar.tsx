import { Progress } from "flowbite-react";
import { isNaN } from "lodash";
import { isPercentageValid } from "../../routes/authenticatedRoutes/flag/helpers/percentage.ts";

export const GoffPercentageProgressBar = ({
  progress,
}: {
  progress: number;
}) => {
  return (
    <div>
      <ProgressBar progress={progress} />
    </div>
  );
};

const ProgressBar = ({ progress }: { progress: number }) => {
  const progressBarColor = () => {
    if (progress > 100) {
      return "red";
    }
    if (progress === 100) {
      return "green";
    }
    return "yellow";
  };

  return (
    <>
      {isNaN(progress) && (
        <span className={"text-sm italic text-red-400"}>invalid</span>
      )}
      {isPercentageValid(progress) && (
        <span className={"text-sm"}>{progress}%</span>
      )}
      {!isNaN(progress) && !isPercentageValid(progress) && (
        <span className={"text-sm italic text-red-400"}>{progress}%</span>
      )}
      <Progress
        progress={progress}
        color={progressBarColor()}
        className={"max-w-xl"}
      />
    </>
  );
};
