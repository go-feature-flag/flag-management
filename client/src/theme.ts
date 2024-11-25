import { type CustomFlowbiteTheme } from "flowbite-react";

export const flowbiteTheme: CustomFlowbiteTheme = {
  footer: {
    root: {
      base: "flex flex-col",
    },
    brand: {
      base: "m-6 flex items-center",
    },
    groupLink: {
      base: "flex flex-col flex-wrap text-gray-500 dark:text-white",
      link: {
        base: "mb-4 last:mr-0 md:mr-6",
      },
    },
    icon: {
      base: "text-gray-400 hover:text-gray-900 dark:hover:text-white",
    },
  },
  modal: {
    body: {
      base: "space-y-6 px-6 pb-4 sm:pb-6 lg:px-8 xl:pb-8",
    },
  },
  sidebar: {
    root: {
      base: "h-full bg-gray-50",
      inner:
        "h-full overflow-y-auto overflow-x-hidden bg-white py-4 px-3 dark:bg-gray-800",
    },
    collapse: {
      list: "space-y-2 py-2 list-none",
    },
    item: {
      base: "no-underline flex items-center rounded-lg p-2 text-lg font-normal text-gray-900 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700",
    },
    itemGroup: {
      base: "list-none border-t border-gray-200 pt-3 first:mt-0 first:border-t-0 first:pt-0 dark:border-gray-700",
    },
  },
  tabs: {
    tablist: {
      base: "flex text-center",
      variant: {
        default: "flex-wrap border-b border-gray-200 dark:border-gray-700",
        underline:
          "flex-wrap -mb-px border-b border-gray-200 dark:border-gray-700",
        pills:
          "flex-wrap font-medium text-sm text-gray-500 dark:text-gray-400 space-x-2",
        fullWidth:
          "w-full text-sm font-medium divide-x divide-gray-200 shadow grid grid-flow-col dark:divide-gray-700 dark:text-gray-400 rounded-none",
      },
      tabitem: {
        base: "flex items-center justify-center p-4 rounded-t-lg text-sm font-medium first:ml-0 disabled:cursor-not-allowed disabled:text-gray-400 disabled:dark:text-gray-500",
        variant: {
          default: {
            base: "rounded-t-lg",
            active: {
              on: "bg-gray-100 text-cyan-600 dark:bg-gray-800 dark:text-cyan-500",
              off: "text-gray-500 hover:bg-gray-50 hover:text-gray-600 dark:text-gray-400 dark:hover:bg-gray-800  dark:hover:text-gray-300",
            },
          },
        },
        icon: "mr-2 h-5 w-5",
      },
    },
  },
  alert: {
    color: {
      failure:
        "mb-4 flex rounded-lg bg-red-50 p-4 text-sm text-red-800 dark:bg-gray-800 dark:text-red-400",
    },
  },
  floatingLabel: {
    input: {
      default: {
        outlined: {
          sm: "dark:bg-gray-800 border-1 peer block w-full appearance-none rounded-lg border-gray-300 bg-transparent px-2.5 pb-2.5 pt-4 text-xs text-gray-900 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500",
          md: "dark:bg-gray-800 border-1 peer block w-full appearance-none rounded-lg border-gray-300 bg-transparent px-2.5 pb-2.5 pt-4 text-sm text-gray-900 focus:border-blue-600 focus:outline-none focus:ring-0 dark:border-gray-600 dark:text-white dark:focus:border-blue-500",
        },
      },
      error: {
        outlined: {
          sm: "dark:bg-gray-800 block px-2.5 pb-2.5 pt-4 w-full text-xs text-gray-900 bg-transparent rounded-lg border-1 appearance-none dark:text-white dark:border-red-500 border-red-600 dark:focus:border-red-500 focus:outline-none focus:ring-0 focus:border-red-600 peer",
          md: "dark:bg-gray-800 block px-2.5 pb-2.5 pt-4 w-full text-sm text-gray-900 bg-transparent rounded-lg border-1 appearance-none dark:text-white dark:border-red-500 border-red-600 dark:focus:border-red-500 focus:outline-none focus:ring-0 focus:border-red-600 peer",
        },
      },
    },
    label: {
      default: {
        outlined: {
          sm: "dark:bg-gray-800 absolute left-1 top-2 z-10 origin-[0] -translate-y-4 scale-75 transition-transform bg-white px-2 text-xs text-gray-500 duration-300 peer-placeholder-shown:top-1/2 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:scale-100 peer-focus:top-2 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:px-2 peer-focus:text-blue-600 dark:text-gray-400 peer-focus:dark:text-blue-500",
          md: "dark:bg-gray-800 absolute left-1 top-2 z-10 origin-[0] -translate-y-4 scale-75 transition-transform bg-white px-2 text-sm text-gray-500 duration-300 peer-placeholder-shown:top-1/2 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:scale-100 peer-focus:top-2 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:px-2 peer-focus:text-blue-600 dark:text-gray-400 peer-focus:dark:text-blue-500",
        },
      },
      success: {
        outlined: {
          sm: "absolute left-1 top-2 z-10 origin-[0] -translate-y-4 scale-75 transition-transform bg-white px-2 text-sm text-green-600 duration-300 peer-placeholder-shown:top-1/2 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:scale-100 peer-focus:top-2 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:px-2 dark:bg-gray-800 dark:text-green-500",
          md: "absolute left-1 top-2 z-10 origin-[0] -translate-y-4 scale-75 transition-transform bg-white px-2 text-sm text-green-600 duration-300 peer-placeholder-shown:top-1/2 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:scale-100 peer-focus:top-2 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:px-2 dark:bg-gray-800 dark:text-green-500",
        },
      },
      error: {
        outlined: {
          sm: "absolute left-1 top-2 z-10 origin-[0] -translate-y-4 scale-75 transition-transform bg-white px-2 text-xs text-red-600 duration-300 peer-placeholder-shown:top-1/2 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:scale-100 peer-focus:top-2 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:px-2 dark:bg-gray-800 dark:text-red-500",
          md: "absolute left-1 top-2 z-10 origin-[0] -translate-y-4 scale-75 transition-transform bg-white px-2 text-xs text-red-600 duration-300 peer-placeholder-shown:top-1/2 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:scale-100 peer-focus:top-2 peer-focus:-translate-y-4 peer-focus:scale-75 peer-focus:px-2 dark:bg-gray-800 dark:text-red-500",
        },
      },
    },
  },
};
