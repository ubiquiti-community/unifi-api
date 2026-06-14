# Unifi UI Crawler

Spin up the Unifi web application using the provided docker-compose.yaml.

Reference the setup code in /home/appkins/src/ubiquiti-community/terraform-provider-unifi/unifi/provider_test.go to write a script that does the same thing before running the crawler.

Instead of running the commands remotely, just click through the UI setup in the hosted network application from the compose.

We need to recursively interact with every Unifi Network Application UI element.

- Traverse the side nav 1 by 1
  - Page loop

  - Click through all buttons
  - Submit all forms (recursively to determine the validation logic) (POST)
  - Submit the form again with changes (UPDATE/PUT/PATCH)
  - Finally, press delete or similar

The script should export a HAR dump of all XHR methods to use for generating OpenAPI spec.

Refer to the following recursive crawler as a starting point:

```
import { Actor } from 'apify';
import { PuppeteerCrawler } from 'crawlee';

await Actor.init();

const crawler = new PuppeteerCrawler({
    async requestHandler({ request, page, enqueueLinks }) {
        const title = await page.title();
        console.log(`Title of ${request.url}: ${title}`);

        await enqueueLinks({
            pseudoUrls: ['https://www.iana.org/[.*]'],
        });
    },
    maxRequestsPerCrawl: 10,
});

await crawler.run(['https://www.iana.org/']);

await Actor.exit();
```

Get creative - we need this to be 100% thorough.