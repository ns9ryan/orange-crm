import qrcode
import argparse
import os
import sys
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad
import base64
from PIL import Image, ImageDraw, ImageFont




def encrypt_aes_cbc(plaintext, key):
    """使用AES-CBC模式加密数据"""
    key_bytes = key.encode('utf-8')
    if len(key_bytes) > 16:
        key_bytes = key_bytes[:16]
    elif len(key_bytes) < 16:
        key_bytes = pad(key_bytes, 16)
    iv = b'\x00' * 16
    cipher = AES.new(key_bytes, AES.MODE_CBC, iv)
    plaintext_bytes = plaintext.encode('utf-8')
    padded_data = pad(plaintext_bytes, 16)
    ciphertext = cipher.encrypt(padded_data)
    return base64.b64encode(ciphertext).decode('utf-8')




def generate_qrcode(content, username, output_path):
    """生成二维码并在下方添加用户名文字，保存到指定路径"""
    try:
        qr = qrcode.QRCode(
            version=1,
            error_correction=qrcode.constants.ERROR_CORRECT_L,
            box_size=10,
            border=4,
        )


        qr.add_data(content)
        qr.make(fit=True)


        img = qr.make_image(fill_color="black", back_color="white").convert('RGB')


        width, height = img.size
        text_height_area = 0  # 文字区域高度


        new_img = Image.new('RGB', (width, height + text_height_area), 'white')
        new_img.paste(img, (0, 0))


        draw = ImageDraw.Draw(new_img)


        try:
            font = ImageFont.truetype("arial.ttf", 20)
        except IOError:
            font = ImageFont.load_default()


        bbox = draw.textbbox((0, 0), username, font=font)
        text_width = bbox[2] - bbox[0]
        text_height = bbox[3] - bbox[1]


        spacing = -20  # 二维码和文字之间的间距


        text_x = (width - text_width) // 2
        text_y = height + spacing + (text_height_area - text_height) // 2


        draw.text((text_x, text_y), username, fill='black', font=font)


        os.makedirs(os.path.dirname(os.path.abspath(output_path)), exist_ok=True)
        new_img.save(output_path)
        print(f"二维码已成功生成并保存到: {output_path}")
        return True
    except Exception as e:
        print(f"生成二维码时出错: {str(e)}")
        return False




def main():
    parser = argparse.ArgumentParser(description='生成用户ID的二维码，支持AES加密，二维码下方显示用户名')


    parser.add_argument('user_id', help='要编码到二维码中的用户ID')
    parser.add_argument('username', help='二维码下方显示的用户名')
    parser.add_argument('enable_encryption', choices=['0', '1'], help='是否开启加密，0表示不加密，1表示加密')
    parser.add_argument('aes_key', help='AES加密密钥（当enable_encryption=1时使用）')
    parser.add_argument('output_path', help='生成的二维码保存路径')


    args = parser.parse_args()


    content = args.user_id
    if args.enable_encryption == '1':
        try:
            content = encrypt_aes_cbc(args.user_id, args.aes_key)
            print(f"已使用AES-CBC模式加密用户ID")
        except Exception as e:
            print(f"加密失败: {str(e)}")
            return 1


    success = generate_qrcode(content, args.username, args.output_path)
    return 0 if success else 1




if __name__ == "__main__":
    sys.exit(main())