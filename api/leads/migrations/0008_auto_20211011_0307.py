# Generated by Django 3.1.13 on 2021-10-11 03:07

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('accounts', '0003_auto_20211010_2310'),
        ('leads', '0007_auto_20211011_0306'),
    ]

    operations = [
        migrations.AlterField(
            model_name='testimony',
            name='member',
            field=models.ForeignKey(null=True, on_delete=django.db.models.deletion.CASCADE, to='accounts.member'),
        ),
    ]