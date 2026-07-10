#!/usr/bin/env node

async function main() {
    const input = await new Promise((resolve, reject) => {
        let data = "";

        process.stdin.setEncoding("utf8");
        process.stdin.on("data", chunk => (data += chunk));
        process.stdin.on("end", () => resolve(data));
        process.stdin.on("error", reject);
    });

    const json = JSON.parse(input);

    const basename = json.entry_info.basename;

    const [id, ...rest] = basename.split("_");

    const output = {
        metas: {
            id: id,
            slug: rest.join("_"),
        },
    }

    console.log(JSON.stringify(output, null, 2));
}

main().catch(err => {
    console.error(err);
    process.exit(1);
});
