// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require("prism-react-renderer/themes/github");
const darkCodeTheme = require("prism-react-renderer/themes/dracula");
const math = require("remark-math");
const katex = require("rehype-katex");

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "Enigma Documentation",
  tagline: "Welcome to Enigma - the blockchain built for finance!",
  url: "https://docs.enigma.network/",
  baseUrl: "/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "throw",
  favicon: "img/favicon.png",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "Enigma", // Usually your GitHub org/user name.
  projectName: "enigma-core/docs", // Usually your repo name.

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  plugins: [],

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          // sidebarPath: require.resolve("./sidebars.js"),
          routeBasePath: "/", // Serve the docs at the site's root
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          //editUrl:
          //  "https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/",
          remarkPlugins: [math],
          rehypePlugins: [katex],
        },
        blog: false,
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      }),
    ],
  ],

  stylesheets: [
    {
      href: "https://cdn.jsdelivr.net/npm/katex@0.13.24/dist/katex.min.css",
      type: "text/css",
      integrity:
        "sha384-odtC+0UGzzFL/6PNoE8rX/SPcQDXBJ+uRepguP4QkPCm2LBxH3FA3y+fKSiJ+AmM",
      crossorigin: "anonymous",
    },
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      prism: {
        additionalLanguages: ['Rust'],
      },
      image: "/img/fury_meta.png",
      navbar: {
        title: "DOCS",
        logo: {
          alt: "Enigma Docs",
          src: "img/enigma.svg",
        },
        items: [
          {
            href: "https://enigma.com/",
            label: "Enigma",
            position: "right",
          },
          {
            href: "https://github.com/EnigmasLab",
            label: "GitHub",
            position: "right",
          },
          {
            type: "search",
            position: "right",
          },
        ],
      },
      footer: {
        style: "light",
        logo: {
          href: "https://enigma.network",
          target: "_blank",
          srcDark: "img/enigma_logo.svg",
          src: "img/enigma_logo_dark.svg",
          height: "36px",
          style: { textAlign: "left" },
          alt: "Enigma Logo",
        },
        links: [
          {
            title: "Enigma",
            items: [
              {
                label: "Hub",
                href: "https://hub.enigma.network/",
              },
              {
                label: "Explorer",
                href: "https://explorer.enigma.network/",
              },
              {
                label: "Blog",
                href: "https://blog.enigma.com/",
              },
            ],
          },
          {
            title: "Community",
            items: [
              {
                label: "Blog",
                href: "https://blog.enigma.com/",
              },
              {
                label: "Enigma Forum",
                href: "https://gov.enigma.network/",
              },
              {
                label: "Discord",
                href: "https://discord.gg/enigma",
              },
              {
                label: "Reddit",
                href: "https://www.reddit.com/r/enigma/",
              },
            ],
          },
          {
            title: "Social",
            items: [
              {
                label: "Twitter",
                href: "https://twitter.com/Enigma_",
              },
              {
                label: "Youtube",
                href: "https://www.youtube.com/channel/UCN99m0dicoMjNmJV9mxioqQ",
              },
              {
                label: "LinkedIn",
                href: "https://www.linkedin.com/company/enigma-protocol",
              },
              {
                label: "Medium",
                href: "https://enigmaprotocol.medium.com/",
              },
            ],
          },
        ],
        copyright: `Copyright © Enigma Labs Inc. since 2021. All rights reserved <a href="https://injectivelabs.org/">Enigma</a>`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
      algolia: {
        appId: "OSH0IFX0OC",
        apiKey: "b8c3930ea2d1ed063992787837d3567f",
        indexName: "enigma",
        contextualSearch: true,
        searchParameters: {},
      },
    }),
};

module.exports = config;
