import os
import json
import xml.etree.ElementTree as ET
import zipfile
import platform

class Example:
    def __init__(self, name, value):
        self.name = name
        self.value = value

def wait_for_enter():
    input("\nPress Enter to continue...")

def list_disks():
    print("Information about logical disks:")
    if platform.system() == "Windows":
        os.system("wmic logicaldisk get caption,volumename,size,filesystem")
    elif platform.system() == "Darwin":  # macOS
        os.system("df -h")
    else:
        print("Function not supported for this operating system")

def create_file(file_name):
    content = input("Enter a string to write to the file: ")
    with open(file_name, 'w') as file:
        file.write(content)
    print(f"File successfully created: {file_name}")

def read_file(file_name):
    with open(file_name, 'r') as file:
        data = file.read()
    print(f"Contents of the file:\n{data}")

def delete_file(file_name):
    try:
        os.remove(file_name)
        print(f"File successfully deleted: {file_name}")
    except Exception as e:
        print(f"Error deleting file: {e}")

def create_json(file_name):
    name = input("Enter name: ")

    while True:
        try:
            value = int(input("Enter value (integer only): "))
            break  # Break the loop if the conversion is successful
        except ValueError:
            print("Invalid input. Please enter an integer.")

    example_json = Example(name, value).__dict__
    
    with open(file_name, 'w') as file:
        json.dump(example_json, file, indent=4)
    print(f"JSON file successfully created: {file_name}")

def read_json(file_name):
    with open(file_name, 'r') as file:
        data = file.read()
    print(f"Contents of JSON file:\n{data}")

def create_xml(file_name):
    name = input("Enter name: ")
    
    while True:
        try:
            value = int(input("Enter value (integer only): "))
            break  # Break the loop if the conversion is successful
        except ValueError:
            print("Invalid input. Please enter an integer.")

    example_xml = Example(name, value)

    root = ET.Element("Example")
    ET.SubElement(root, "name").text = example_xml.name
    ET.SubElement(root, "value").text = str(example_xml.value)

    tree = ET.ElementTree(root)
    tree.write(file_name, xml_declaration=True, encoding='utf-8')
    print(f"XML file successfully created: {file_name}")

def read_xml(file_name):
    tree = ET.parse(file_name)
    root = tree.getroot()
    print(f"Contents of XML file:\n{ET.tostring(root, encoding='utf-8').decode('utf-8')}")

def create_zip_archive(zip_file_name, file_name):
    with zipfile.ZipFile(zip_file_name, 'w') as zipf:
        zipf.write(file_name, os.path.basename(file_name))
    print(f"ZIP archive successfully created: {zip_file_name}")

def unzip_archive(zip_file_name, dest_dir):
    with zipfile.ZipFile(zip_file_name, 'r') as zipf:
        zipf.extractall(dest_dir)
    print(f"Files successfully extracted to: {dest_dir}")

def main():
    # List disks (only for Windows and macOS)
    list_disks()
    wait_for_enter()

    # File operations
    print("\nFile operations:")
    file_name = "example.txt"
    create_file(file_name)
    wait_for_enter()

    read_file(file_name)
    wait_for_enter()

    delete_file(file_name)
    wait_for_enter()

    # JSON operations
    print("\nJSON operations:")
    json_file_name = "example.json"
    create_json(json_file_name)
    wait_for_enter()

    read_json(json_file_name)
    wait_for_enter()

    delete_file(json_file_name)
    wait_for_enter()

    # XML operations
    print("\nXML operations:")
    xml_file_name = "example.xml"
    create_xml(xml_file_name)
    wait_for_enter()

    read_xml(xml_file_name)
    wait_for_enter()

    delete_file(xml_file_name)
    wait_for_enter()

    # ZIP archive operations
    print("\nZIP archive operations:")
    file_to_archive = "file_for_zip.txt"
    create_file(file_to_archive)
    wait_for_enter()

    zip_file_name = "archive.zip"
    create_zip_archive(zip_file_name, file_to_archive)
    wait_for_enter()

    unzip_archive(zip_file_name, "./unzipped")
    wait_for_enter()

    delete_file(file_to_archive)
    delete_file(zip_file_name)
    wait_for_enter()

if __name__ == "__main__":
    main()