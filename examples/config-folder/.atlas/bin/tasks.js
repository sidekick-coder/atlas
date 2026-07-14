#!/usr/bin/env node

import fs from "fs";
import path from "path";

const OUTPUT_DIR = path.resolve(import.meta.dirname, "..", "..", "tasks");
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

const tagPool = [
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
];

const verbs = [
    "Implement",
    "Fix",
    "Refactor",
    "Review",
    "Optimize",
    "Design",
    "Document",
    "Test",
    "Deploy",
    "Investigate",
    "Build",
    "Improve",
];

const adjectives = [
    "secure",
    "scalable",
    "distributed",
    "responsive",
    "efficient",
    "modular",
    "reliable",
    "automated",
    "dynamic",
    "optimized",
];

const nouns = [
    "API",
    "authentication",
    "dashboard",
    "database",
    "pipeline",
    "cache",
    "worker",
    "frontend",
    "backend",
    "CLI",
    "scheduler",
    "search",
];

const firstNames = [
    "Alice",
    "Bob",
    "Charlie",
    "Diana",
    "Emma",
    "Lucas",
    "Noah",
    "Olivia",
    "Sophia",
    "Liam",
];

const lastNames = [
    "Smith",
    "Johnson",
    "Brown",
    "Miller",
    "Wilson",
    "Taylor",
    "Davis",
    "Moore",
    "Clark",
    "Walker",
];

const paragraphs = [
    "Review the implementation and document any important findings before deployment.",
    "Coordinate with the team to validate requirements and verify the expected behavior.",
    "Monitor performance after release and collect feedback for future improvements.",
    "Ensure all tests pass and update any relevant documentation before closing the task.",
    "Investigate edge cases and confirm the solution works across supported environments.",
];

function sample(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}

function randInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function randBool() {
    return Math.random() < 0.5;
}

function sampleMany(arr, min, max) {
    const copy = [...arr];
    copy.sort(() => Math.random() - 0.5);
    return copy.slice(0, randInt(min, max));
}

function randomDateBetween(from, to) {
    return new Date(
        from.getTime() + Math.random() * (to.getTime() - from.getTime())
    );
}

function randomPastDate(years = 2) {
    const now = new Date();
    const from = new Date(now);
    from.setFullYear(now.getFullYear() - years);
    return randomDateBetween(from, now);
}

function randomFutureDate(days = 120) {
    const now = new Date();
    const to = new Date(now);
    to.setDate(now.getDate() + days);
    return randomDateBetween(now, to);
}

function randomName() {
    return `${sample(firstNames)} ${sample(lastNames)}`;
}

function randomTitle() {
    return `${sample(verbs)} ${sample(adjectives)} ${sample(nouns)}`;
}

function randomParagraphs() {
    return Array.from(
        { length: randInt(5, 10) },
        () => sample(paragraphs)
    ).join("\n\n");
}

for (let i = 1; i <= TOTAL_TASKS; i++) {
    const number = String(i).padStart(3, "0");

    const title = randomTitle();

    const slug = title
        .toLowerCase()
        .replace(/\s+/g, "_")
        .replace(/[^\w-]/g, "");

    const filename = `${number}_${slug}.md`;

    const created = randomPastDate(2);
    const updated = randomDateBetween(created, new Date());

    const due =
        Math.random() > 0.35
            ? randomFutureDate(120)
            : null;

    const completed =
        Math.random() > 0.75
            ? randomDateBetween(updated, new Date())
            : null;

    const status = completed ? "done" : sample(statuses);

    const priority = sample(priorities);

    const tags = sampleMany(tagPool, 1, 5);

    const content = `---
id: task-${number}
title: "${title}"
status: ${status}
priority: ${priority}
project: ${sample(projects)}
context: ${sample(contexts)}
favorite: ${randBool()}
archived: ${randBool()}
estimateHours: ${randInt(1, 40)}
progress: ${randInt(0, 100)}
tags:
${tags.map((t) => `  - ${t}`).join("\n")}
created: ${created.toISOString()}
updated: ${updated.toISOString()}
due: ${due ? due.toISOString() : "null"}
completed: ${completed ? completed.toISOString() : "null"}
assignee: ${randomName()}
reviewer: ${randomName()}
source: generated
---

# ${title}

- [ ] ${sample(verbs)} ${sample(nouns)}
- [ ] ${sample(verbs)} ${sample(nouns)}
- [ ] ${sample(verbs)} ${sample(nouns)}

## Notes

${randomParagraphs()}
`;

    fs.writeFileSync(path.join(OUTPUT_DIR, filename), content);
}

console.log(`Generated ${TOTAL_TASKS} task files in ${OUTPUT_DIR}`);
