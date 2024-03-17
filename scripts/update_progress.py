import re
from html.parser import HTMLParser

from markdown_it import MarkdownIt

readme_path = "README.md"


class HTMLTaskParser(HTMLParser):

    def __init__(self):
        super().__init__()
        self.reset()
        self.solved = 0
        self.total = 0
        self.in_li = False
        self.li_data = []

    def handle_starttag(self, tag, attrs):
        if tag == 'li':
            self.in_li = True

    def handle_endtag(self, tag):
        if tag == 'li':
            self.in_li = False

    def handle_data(self, data):
        if self.in_li:
            data = data.strip()
            if data == '[ ]':
                self.total += 1
            if data == '[x]':
                self.total += 1
                self.solved += 1

    def parse_tasks(self, html: str):
        self.feed(html)


def calculate_progress():
    with open(readme_path, "r+", encoding='utf-8') as file:
        content = file.read()
        md = MarkdownIt()
        html = md.render(content)
        parser = HTMLTaskParser()
        parser.parse_tasks(html)
        return int((parser.solved / parser.total) * 100)


def update_readme(progress):
    with open(readme_path, "r+", encoding="utf-8") as file:
        content = file.read()
        new_content = re.sub(r"\bprogress-\d+%", f"progress-{progress}%", content)
        file.seek(0)
        file.write(new_content)
        file.truncate()


progress = calculate_progress()
update_readme(progress)
