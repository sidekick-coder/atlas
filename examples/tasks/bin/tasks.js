import fs from "fs";
import path from "path";
import { faker } from "@faker-js/faker";

const OUTPUT_DIR = path.resolve(__dirname, "tasks");
const TOTAL_TASKS = Number(process.argv[2]) || 10;

fs.rmSync(OUTPUT_DIR, { recursive: true, force: true });

fs.mkdirSync(OUTPUT_DIR, { recursive: true });

const priorities = ["low", "medium", "high", "critical"];
const statuses = [
  "todo",
  "in-progress",
  "blocked",
  "done",
  "cancelled",
];

const contexts = [
  "work",
  "personal",
  "home",
  "study",
  "health",
  "finance",
];

const projects = [
  "Alpha",
  "Bravo",
  "Charlie",
  "Delta",
  "Echo",
  "Phoenix",
  "Atlas",
];

for (let i = 1; i <= TOTAL_TASKS; i++) {
  const number = String(i).padStart(3, "0");

  const title = faker.hacker.phrase().replace(/[^\w\s-]/g, "");
  const slug = title
    .toLowerCase()
    .replace(/\s+/g, "_")
    .replace(/[^\w-]/g, "");

  const filename = `${number}_${slug}.md`;

  const created = faker.date.past({ years: 2 });
  const updated = faker.date.between({
    from: created,
    to: new Date(),
  });

  const due =
    Math.random() > 0.35
      ? faker.date.soon({ days: 120 })
      : null;

  const completed =
    Math.random() > 0.75
      ? faker.date.between({
          from: updated,
          to: new Date(),
        })
      : null;

  const status = completed
    ? "done"
    : faker.helpers.arrayElement(statuses);

  const priority = faker.helpers.arrayElement(priorities);

  const tags = faker.helpers.arrayElements(
    [
      "backend",
      "frontend",
      "bug",
      "feature",
      "documentation",
      "urgent",
      "meeting",
      "research",
      "testing",
      "design",
      "refactor",
      "personal",
      "api",
      "mobile",
    ],
    faker.number.int({ min: 1, max: 5 })
  );

  const content = `---
id: task-${number}
title: "${title}"
status: ${status}
priority: ${priority}
project: ${faker.helpers.arrayElement(projects)}
context: ${faker.helpers.arrayElement(contexts)}
favorite: ${faker.datatype.boolean()}
archived: ${faker.datatype.boolean()}
estimateHours: ${faker.number.int({ min: 1, max: 40 })}
progress: ${faker.number.int({ min: 0, max: 100 })}
tags:
${tags.map((t) => `  - ${t}`).join("\n")}
created: ${created.toISOString()}
updated: ${updated.toISOString()}
due: ${due ? due.toISOString() : "null"}
completed: ${completed ? completed.toISOString() : "null"}
assignee: ${faker.person.fullName()}
reviewer: ${faker.person.fullName()}
source: generated
---

# ${title}

- [ ] ${faker.hacker.verb()} ${faker.hacker.noun()}
- [ ] ${faker.hacker.verb()} ${faker.hacker.noun()}
- [ ] ${faker.hacker.verb()} ${faker.hacker.noun()}

## Notes

${faker.lorem.paragraphs({ min: 2, max: 5 })}
`;

  fs.writeFileSync(path.join(OUTPUT_DIR, filename), content);
}

console.log(`Generated ${TOTAL_TASKS} task files in ${OUTPUT_DIR}`);
