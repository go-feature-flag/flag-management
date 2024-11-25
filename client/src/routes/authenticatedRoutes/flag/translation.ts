export const flagPageTranslation = {
  en: {
    translation: {
      component: {
        exitWithoutSaveModal: {
          text: "Are you sure you want to exit without saving?",
        },
        percentageRollout: {
          title: "💯 Percentage Rollout",
          description:
            "A percentage rollout means that your users are divided in different buckets and you serve different variations to them. " +
            "Note that an evaluation context will always have the same variation.",
          alertVariation: "Please specify at least one variation with a name.",
          errors: {
            variationErrorTitle: "Percentage Error",
            required: "Please provide a percentage for the variation",
            invalidPercentage: "Percentage should be between 0% and 100%",
            invalidPercentageSum: "The sum of all percentages should be 100%",
          },
        },
        progressiveRollout: {
          title: "📈 Progressive Rollout",
          description:
            "A progressive rollout allows you to increase the percentage of your flag " +
            "over time. Configure your rollout and we will release your flag progressively.",
          step1: {
            part1: "Progressively transition from",
            part2: "to",
          },
          step2: {
            part1: "Over",
            part2: "TO",
            part3: "time range",
          },
          step3: {
            part1: "Starting at",
            part2: "and stopping at",
          },
          errors: {
            percentageFromRequired: "Please provide a START percentage",
            percentageToRequired: "Please provide a STOP percentage",
            percentageToHigherThanFrom:
              "The to percentage must be higher than the from percentage",
            variationFromRequired:
              "Please select a variation for the from field",
            variationToRequired: "Please select a variation for the to field",
            percentageErrorTitle: "Percentage Error",
            variationErrorTitle: "Variation Error",
            invalidPercentageFrom:
              "Start percentage should be between 0% and 100%",
            invalidPercentageTo: "End percentage should be between 0% and 100%",
          },
        },
        variationSelector: {
          serve: "Serve",
          selectRollout: {
            percentage: "💯 a percentage rollout",
            progressive: "📈 a progressive rollout",
            errors: {
              noSelectedVariation: "Please select a variation",
            },
          },
        },
      },
      page: {
        flag: {
          pageTitleNew: "New Feature Flag",
          pageTitleEdit: "Edit Feature Flag",
          copyTooltip: "Copy Feature Flag name",
          metadata: {
            title: "Metadata",
            removeButton: "Remove Metadata",
            description:
              "By adding metadata to a feature flag, you can enrich your feature flag " +
              "with information, such as a configuration URL or the " +
              "originating issue tracker link ...",
            messageNoMetadata:
              "Your feature flag has no metadata, start adding some!",
            fields: {
              name: {
                title: "Name",
                errorRequired: "Please provide a name for your metadata",
              },
              value: {
                title: "Value",
                errorRequired: "Please provide a value for your metadata",
              },
            },
          },
          settings: {
            trackEvent: {
              title: "Track Events",
              description:
                "Disable if you don't want to track events for this feature flag.",
            },
            description: {
              title: "Description",
              description: "Add a description to your feature flag.",
            },
            version: {
              title: "Version",
              description: "Add a version to your feature flag.",
            },
            normalZone: "Settings",
            dangerZone: "Danger Zone",
            delete: {
              error:
                "Error while trying to delete your feature flag. {{error}}",
              title: "Delete Feature Flag",
              description:
                "This will permanently delete the feature flag and cannot be undone.",
              button: "Delete",
              modal:
                "Are you sure you want to delete the Feature Flag named {{name}}?",
            },
            status: {
              error:
                "Error while trying to update your feature flag status. {{error}}",
              enable: {
                title: "Enable Feature Flag",
                description:
                  "Enable the feature flag and it will be available to be user everywhere.",
                button: "Enable",
                modal:
                  "Are you sure you want to enable the Feature Flag named {{name}}?",
              },
              disable: {
                title: "Disable Feature Flag",
                description:
                  "Disable the feature flag and we will not evaluate it for any users.",
                button: "Disable",
                modal:
                  "Are you sure you want to disable the Feature Flag named {{name}}?",
              },
            },
          },
          addButton: "Add Variation",
          variations: {
            defaultRule: {
              info: "When this feature flag is enable and contexts don't match any targeting rules.",
            },
            section: {
              defaultRule: "Default Rule",
              variations: "Variations",
            },
            tooltip: {
              removeOk: "Remove Variation",
              removeMinVariations: "You need at least 2 variations",
              removeVariationUsedDefaultRule:
                "This variation is used in the default rule",
              removeVariationUsedTargetingRule:
                "This variation is used in a targeting rule",
            },
            fields: {
              name: "Name",
              value: "Value",
              error: {
                nameRequired: "Please provide a name for the variation",
                valueRequired: "Please provide a value for the variation",
              },
            },
          },
          targeting: {
            title: "Targeting",
            description:
              "By creating rules, you enable some capabilities to a segment of your users. This is ideal if you want to test in production, or enable a feature for a subset of your users.",
            addRuleTopButton: {
              title: "Add Rule",
              notice:
                "Rule evaluation runs top to bottom, and serves a variation when a target matches",
            },
            rule: {
              dropdown: {
                advanced: "Advanced Query",
                queryBuilder: "Query Builder",
                delete: "Delete Rule",
                disable: "Disable Rule",
                enable: "Enable Rule",
              },
              deleteConfirmation:
                'Are you sure you want to delete the targeting rule named "{{name}}"?',
            },
            queryBuilder: {
              addRule: "+ Add Condition",
              addGroup: "+ Add Group",
            },
            advancedRule: {
              queryFormat: {
                title: "Rule Format",
                tableColumn: {
                  exp: "Expression",
                  mean: "Meaning",
                },
                description1:
                  "The rule format is based on the nikunjy/rules library.",
                description2:
                  "All the operations can be written in capitalized or lowercase characters (ex: eq or EQ can be used).",
                description3:
                  "Logical Operations supported are AND & OR. Compare Expression and their definitions (a|b means you can use one of either a or b):",
              },
              errors: {
                queryRequired:
                  'Please provide a query for your targeting rule named "{{name}}"',
              },
            },
          },
          error: {
            dataFetch: "Error while loading your feature flag information.",
          },
          ruleOperators: {
            eq: "Equal",
            ne: "Not Equals To",
            lt: "Less Than",
            gt: "Greater Than",
            le: "Less Than or Equal",
            ge: "Greater Than or Equal",
            co: "Contains",
            sw: "Starts With",
            ew: "Ends With",
            in: "In a List",
            pr: "Present",
            not: "Not",
          },
        },
      },
    },
  },
};
