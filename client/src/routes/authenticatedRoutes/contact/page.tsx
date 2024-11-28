import { Card } from "flowbite-react";
import { t } from "i18next";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import community from "../../../assets/pages/contact/community.svg";
import contact from "../../../assets/pages/contact/contact.svg";
import support from "../../../assets/pages/contact/support.svg";
import { useDocumentTitle } from "../../../hooks/documentTitle.ts";

const translationBaseKey = "page.contact";

interface ContactCardProps {
  title: string;
  description: string;
  link: string;
  img: string;
  imgAlt: string;
}

const contactCards: ContactCardProps[] = [
  {
    img: contact,
    imgAlt: "contact us",
    title: t(`${translationBaseKey}.cards.contact.title`),
    description: t(`${translationBaseKey}.cards.contact.description`),
    link: "mailto:contact@gofeatureflag.org",
  },
  {
    img: support,
    imgAlt: "Support",
    title: t(`${translationBaseKey}.cards.support.title`),
    description: t(`${translationBaseKey}.cards.support.description`),
    link: "https://github.com/thomaspoignant/go-feature-flag/issues",
  },
  {
    title: t(`${translationBaseKey}.cards.community.title`),
    img: community,
    imgAlt: "Open-source community",
    description: t(`${translationBaseKey}.cards.community.description`),
    link: "https://gofeatureflag.org/slack",
  },
];

function ContactCard(props: ContactCardProps) {
  return (
    <div className="flex h-full min-h-fit content-center px-2 py-3">
      <Link to={props.link} target="_blank">
        <Card
          className=" h-full max-w-lg place-items-center hover:bg-gray-100 hover:shadow-xl dark:hover:bg-gray-700"
          renderImage={() => (
            <img src={props.img} alt={props.imgAlt} height={190} />
          )}
        >
          <div className={"min-h-fit"}>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {props.title}
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              {props.description}
            </p>
          </div>
        </Card>
      </Link>
    </div>
  );
}

export default function ContactPage() {
  useDocumentTitle("Contact");
  const { t } = useTranslation();
  return (
    <div className="p-3 md:p-6 lg:p-10">
      <section>
        <div className="grid w-full grid-cols-1 place-items-center md:grid-cols-2 lg:grid-cols-3 lg:px-20 ">
          <div className="text-center md:col-span-2 lg:col-span-3">
            <h1 className="py-10 text-7xl tracking-tight dark:text-white">
              {t(`${translationBaseKey}.title`)}
            </h1>
            <div className="mb-5 content-center text-center text-2xl dark:text-white">
              {t(`${translationBaseKey}.subtitle`)}
            </div>
          </div>
          {contactCards.map((card) => (
            <ContactCard
              key={card.title}
              img={card.img}
              imgAlt={card.imgAlt}
              title={card.title}
              description={card.description}
              link={card.link}
            />
          ))}
        </div>
      </section>
    </div>
  );
}
