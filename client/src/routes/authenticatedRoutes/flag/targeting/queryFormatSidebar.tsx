import { Table, Tooltip } from "flowbite-react";
import { useTranslation } from "react-i18next";
import { config } from "../../../../config.ts";

const translationBaseKey = "page.flag.targeting.advancedRule.queryFormat";

export const QueryFormatSidebar = () => {
  const { t } = useTranslation();
  return (
    <>
      <p className={"py-2"}>{t(`${translationBaseKey}.description1`)}</p>
      <p className={"py-2"}>{t(`${translationBaseKey}.description2`)}</p>
      <p className={"py-2"}>{t(`${translationBaseKey}.description3`)}</p>
      <Table striped>
        <Table.Head>
          <Table.HeadCell>
            <Tooltip content={t(`${translationBaseKey}.tableColumn.exp`)}>
              {t(`${translationBaseKey}.tableColumn.exp`)
                .toUpperCase()
                .slice(0, 3)}
            </Tooltip>
          </Table.HeadCell>
          <Table.HeadCell>
            <Tooltip content={t(`${translationBaseKey}.tableColumn.mean`)}>
              {t(`${translationBaseKey}.tableColumn.mean`)}
            </Tooltip>
          </Table.HeadCell>
        </Table.Head>
        <Table.Body className="divide-y">
          {config.ruleOperators.map((operator) => (
            <Table.Row
              key={operator.name}
              className="bg-white dark:border-gray-700 dark:bg-gray-800"
            >
              <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white">
                <div className="relative">
                  <span className={"font-mono"}>
                    {operator.name.toUpperCase()}
                  </span>
                </div>
              </Table.Cell>
              <Table.Cell>{t(operator.translationKey)}</Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </>
  );
};
