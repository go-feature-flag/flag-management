export const flagsPageTranslation = {
  en: {
    translation: {
      page: {
        flags: {
          title: "Feature Flags",
          newModal: {
            createButton: "Create",
            title: "Create a new feature flag 🚀",
            labelNewFlag: "Flag name:",
            labelSelectType: "Type of your flag:",
            infoFlagName:
              "This name will be used by SDKs to retrieve the feature value and enabled state. <1/> This cannot be edited once the feature flag has been created.",
            infoFlagType:
              "The feature flag type cannot be edited once the feature flag has been created.",

            error: {
              required: "Flag name is required",
              exists: "Flag name already exists",
            },
          },
          subtitle: {
            line1: "Manage all your feature flags.",
            line2: "Select a flag to manage the targeting and rollout rules.",
          },
          flagList: {
            noResult: "No result",
            errors: {
              loading: "Error while loading your feature flags list.",
            },
            topBar: {
              searchPlaceholder: "Search",
              createFlagButton: "Create flag",
            },
            row: {
              errors: {
                statusChange:
                  "Error while trying to update your feature flag status.",
                delete: "Error while trying to delete your feature flag.",
              },
              tooltip: {
                created: "Created",
                status: "Flag status",
                info: "Info",
                delete: "Delete",
              },
              modal: {
                delete:
                  "Are you sure you want to delete the Feature Flag named {{name}}?",
                enable:
                  "Are you sure you want to {{action}} the Feature Flag named {{name}}?",
              },
              lastUpdated: "Last Updated",
            },
          },
        },
      },
    },
  },
};
